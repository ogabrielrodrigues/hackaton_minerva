import type { Metadata } from "next"
import { Montserrat } from "next/font/google"
import "./globals.css"
import { Header } from "@/components/header"

const sans = Montserrat({ subsets: ["latin"], variable: "--font-sans" })

export const metadata: Metadata = {
  title: "Hackaton Minerva",
  description: "Hackaton Minerva",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body
        className={`${sans.variable} h-screen w-screen overflow-hidden font-sans antialiased`}
      >
        <Header />
        {children}
      </body>
    </html>
  )
}
