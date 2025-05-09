import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'Monorepo Frontend',
  description: 'Next.js frontend for the monorepo project',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}