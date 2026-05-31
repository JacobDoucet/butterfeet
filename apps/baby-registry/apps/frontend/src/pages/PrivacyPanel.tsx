import type { Registry } from '../api';
import ShippingSection from '../components/privacy/ShippingSection';
import AccessControlSection from '../components/privacy/AccessControlSection';

export default function PrivacyPanel({
  reg,
  section,
}: {
  reg: Registry;
  section: 'shipping' | 'access';
}) {
  if (section === 'shipping') return <ShippingSection reg={reg} />;
  return <AccessControlSection reg={reg} />;
}
