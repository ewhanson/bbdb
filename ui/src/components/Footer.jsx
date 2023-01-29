import { DateTime } from "luxon";

export function Footer() {
  const year = DateTime.now().year;
  const version = APP_VERSION;
  return (
    <footer className="footer footer-center p-4 bg-base-100 text-base-content">
      <div>
        <p>
          Copyright © {year} - Erik Hanson - v{version}
        </p>
      </div>
    </footer>
  );
}
