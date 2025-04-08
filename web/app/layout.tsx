import type { Metadata } from "next";
import { Geist_Mono as GeistMono } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";

const font_mono = GeistMono({
  variable: "--font-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: {
    template: "Imobiliary | %s",
    default: 'Imobiliary',
  }
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body
        className={cn(font_mono.variable, 'font-display antialiased')}
      >
        {children}
      </body>
    </html>
  );
}
