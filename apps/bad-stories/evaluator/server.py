#!/usr/bin/env python3
"""Tiny dev server for the Bad Stories evaluator.

Serves the apps/bad-stories/ tree as static files and exposes a small
feedback API that persists notes per story to apps/bad-stories/feedback/
(gitignored).

Endpoints:
  GET  /api/feedback                -> entire feedback store
  POST /api/feedback                -> {storyKey, text} append entry
  POST /api/feedback/resolve        -> {storyKey, entryId} mark resolved
  POST /api/feedback/delete         -> {storyKey, entryId} remove entry
"""
from __future__ import annotations

import json
import os
import sys
import threading
import uuid
from datetime import datetime, timezone
from http.server import SimpleHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent  # apps/bad-stories/
FEEDBACK_DIR = ROOT / "feedback"
FEEDBACK_FILE = FEEDBACK_DIR / "feedback.json"
_lock = threading.Lock()


def load_store() -> dict:
    if not FEEDBACK_FILE.exists():
        return {}
    try:
        return json.loads(FEEDBACK_FILE.read_text() or "{}")
    except json.JSONDecodeError:
        return {}


def save_store(store: dict) -> None:
    FEEDBACK_DIR.mkdir(parents=True, exist_ok=True)
    tmp = FEEDBACK_FILE.with_suffix(".json.tmp")
    tmp.write_text(json.dumps(store, indent=2, sort_keys=True))
    tmp.replace(FEEDBACK_FILE)


class Handler(SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=str(ROOT), **kwargs)

    def log_message(self, fmt, *args):
        sys.stderr.write("[eval] " + fmt % args + "\n")

    def _json(self, status: int, payload) -> None:
        body = json.dumps(payload).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.send_header("Cache-Control", "no-store")
        self.end_headers()
        self.wfile.write(body)

    def _read_json(self):
        length = int(self.headers.get("Content-Length") or 0)
        raw = self.rfile.read(length) if length else b""
        try:
            return json.loads(raw or b"{}")
        except json.JSONDecodeError:
            return None

    def do_GET(self):
        if self.path == "/api/feedback":
            with _lock:
                self._json(200, load_store())
            return
        return super().do_GET()

    def do_POST(self):
        if self.path == "/api/feedback":
            data = self._read_json() or {}
            key = (data.get("storyKey") or "").strip()
            text = (data.get("text") or "").strip()
            if not key or not text:
                self._json(400, {"error": "storyKey and text required"})
                return
            entry = {
                "id": uuid.uuid4().hex[:12],
                "ts": datetime.now(timezone.utc).isoformat(timespec="seconds"),
                "text": text,
                "status": "open",
            }
            with _lock:
                store = load_store()
                store.setdefault(key, []).append(entry)
                save_store(store)
            self._json(200, {"ok": True, "entry": entry})
            return
        if self.path in ("/api/feedback/resolve", "/api/feedback/delete"):
            data = self._read_json() or {}
            key = (data.get("storyKey") or "").strip()
            entry_id = (data.get("entryId") or "").strip()
            if not key or not entry_id:
                self._json(400, {"error": "storyKey and entryId required"})
                return
            with _lock:
                store = load_store()
                entries = store.get(key, [])
                if self.path.endswith("/delete"):
                    new = [e for e in entries if e.get("id") != entry_id]
                    if len(new) == len(entries):
                        self._json(404, {"error": "not found"})
                        return
                    if new:
                        store[key] = new
                    else:
                        store.pop(key, None)
                else:
                    found = False
                    for e in entries:
                        if e.get("id") == entry_id:
                            e["status"] = "resolved"
                            e["resolved_ts"] = datetime.now(timezone.utc).isoformat(timespec="seconds")
                            found = True
                            break
                    if not found:
                        self._json(404, {"error": "not found"})
                        return
                save_store(store)
            self._json(200, {"ok": True})
            return
        self._json(404, {"error": "not found"})


def main():
    port = int(os.environ.get("PORT", "8765"))
    server = ThreadingHTTPServer(("127.0.0.1", port), Handler)
    print(f"Bad Stories evaluator on http://localhost:{port}/evaluator/")
    print(f"Feedback file: {FEEDBACK_FILE}")
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        pass
    finally:
        server.server_close()


if __name__ == "__main__":
    main()
