// Tiny dev server for the Bad Stories evaluator.
//
// Serves the apps/bad-stories/ tree as static files and exposes a small
// feedback API that persists notes per story to apps/bad-stories/feedback/
// (gitignored).
//
// Endpoints:
//   GET  /api/feedback           -> entire feedback store
//   POST /api/feedback           -> {storyKey, text} append entry
//   POST /api/feedback/resolve   -> {storyKey, entryId} mark resolved
//   POST /api/feedback/delete    -> {storyKey, entryId} remove entry
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type Entry struct {
	ID          string `json:"id"`
	TS          string `json:"ts"`
	Text        string `json:"text"`
	Status      string `json:"status"`
	ResolvedTS  string `json:"resolved_ts,omitempty"`
}

type Store struct {
	mu   sync.Mutex
	path string
}

func (s *Store) load() (map[string][]Entry, error) {
	data, err := os.ReadFile(s.path)
	if errors.Is(err, os.ErrNotExist) || len(data) == 0 {
		return map[string][]Entry{}, nil
	}
	if err != nil {
		return nil, err
	}
	out := map[string][]Entry{}
	if err := json.Unmarshal(data, &out); err != nil {
		return map[string][]Entry{}, nil
	}
	return out, nil
}

func (s *Store) save(m map[string][]Entry) error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	// Stable key order for diffability.
	type kv struct {
		k string
		v []Entry
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ordered := make([]kv, 0, len(keys))
	for _, k := range keys {
		ordered = append(ordered, kv{k, m[k]})
	}
	var buf strings.Builder
	buf.WriteString("{\n")
	for i, e := range ordered {
		kj, _ := json.Marshal(e.k)
		buf.WriteString("  ")
		buf.Write(kj)
		buf.WriteString(": ")
		vj, _ := json.MarshalIndent(e.v, "  ", "  ")
		buf.Write(vj)
		if i < len(ordered)-1 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString("}\n")
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, []byte(buf.String()), 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	body, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func readJSON(r *http.Request, v any) error {
	defer r.Body.Close()
	b, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, v)
}

func newID() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func nowUTC() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8765"
	}

	// Resolve apps/bad-stories/ relative to this source file's directory
	// so the binary works regardless of cwd.
	_, srcFile, _, _ := runtime.Caller(0)
	root := filepath.Clean(filepath.Join(filepath.Dir(srcFile), ".."))
	if abs, err := filepath.Abs(root); err == nil {
		root = abs
	}
	feedbackFile := filepath.Join(root, "feedback", "feedback.json")
	store := &Store{path: feedbackFile}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/feedback", func(w http.ResponseWriter, r *http.Request) {
		store.mu.Lock()
		defer store.mu.Unlock()
		switch r.Method {
		case http.MethodGet:
			m, err := store.load()
			if err != nil {
				writeJSON(w, 500, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, 200, m)
		case http.MethodPost:
			var req struct {
				StoryKey string `json:"storyKey"`
				Text     string `json:"text"`
			}
			if err := readJSON(r, &req); err != nil {
				writeJSON(w, 400, map[string]string{"error": "invalid json"})
				return
			}
			req.StoryKey = strings.TrimSpace(req.StoryKey)
			req.Text = strings.TrimSpace(req.Text)
			if req.StoryKey == "" || req.Text == "" {
				writeJSON(w, 400, map[string]string{"error": "storyKey and text required"})
				return
			}
			m, _ := store.load()
			entry := Entry{ID: newID(), TS: nowUTC(), Text: req.Text, Status: "open"}
			m[req.StoryKey] = append(m[req.StoryKey], entry)
			if err := store.save(m); err != nil {
				writeJSON(w, 500, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, 200, map[string]any{"ok": true, "entry": entry})
		default:
			w.Header().Set("Allow", "GET, POST")
			writeJSON(w, 405, map[string]string{"error": "method not allowed"})
		}
	})

	mutate := func(action string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.Header().Set("Allow", "POST")
				writeJSON(w, 405, map[string]string{"error": "method not allowed"})
				return
			}
			var req struct {
				StoryKey string `json:"storyKey"`
				EntryID  string `json:"entryId"`
			}
			if err := readJSON(r, &req); err != nil {
				writeJSON(w, 400, map[string]string{"error": "invalid json"})
				return
			}
			req.StoryKey = strings.TrimSpace(req.StoryKey)
			req.EntryID = strings.TrimSpace(req.EntryID)
			if req.StoryKey == "" || req.EntryID == "" {
				writeJSON(w, 400, map[string]string{"error": "storyKey and entryId required"})
				return
			}
			store.mu.Lock()
			defer store.mu.Unlock()
			m, _ := store.load()
			entries := m[req.StoryKey]
			switch action {
			case "delete":
				kept := entries[:0:0]
				found := false
				for _, e := range entries {
					if e.ID == req.EntryID {
						found = true
						continue
					}
					kept = append(kept, e)
				}
				if !found {
					writeJSON(w, 404, map[string]string{"error": "not found"})
					return
				}
				if len(kept) == 0 {
					delete(m, req.StoryKey)
				} else {
					m[req.StoryKey] = kept
				}
			case "resolve":
				found := false
				for i := range entries {
					if entries[i].ID == req.EntryID {
						entries[i].Status = "resolved"
						entries[i].ResolvedTS = nowUTC()
						found = true
						break
					}
				}
				if !found {
					writeJSON(w, 404, map[string]string{"error": "not found"})
					return
				}
				m[req.StoryKey] = entries
			}
			if err := store.save(m); err != nil {
				writeJSON(w, 500, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, 200, map[string]bool{"ok": true})
		}
	}
	mux.HandleFunc("/api/feedback/resolve", mutate("resolve"))
	mux.HandleFunc("/api/feedback/delete", mutate("delete"))

	// Static file server for the bad-stories tree.
	fs := http.FileServer(http.Dir(root))
	mux.Handle("/", fs)

	addr := "127.0.0.1:" + port
	log.Printf("Bad Stories evaluator on http://localhost:%s/evaluator/", port)
	log.Printf("Feedback file: %s", feedbackFile)
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
