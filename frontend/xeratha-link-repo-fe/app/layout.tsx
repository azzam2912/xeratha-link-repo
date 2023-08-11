import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Xerahta Link Repo',
  description: 'A bunch of useful links here!',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
