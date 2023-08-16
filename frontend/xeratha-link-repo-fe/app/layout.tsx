import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Xeratha Link Repo',
  description: 'A bunch of useful links here!',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head>
        <link rel='shortcut icon' href='/logo-red-blue.png' />
      </head>
      <body>{children}</body>
    </html>
  )
}
