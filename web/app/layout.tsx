import { Toaster } from "@/components/ui/sonner";
import { cn } from "@/lib/utils";
import type { Metadata } from "next";
import { Geist_Mono as GeistMono } from "next/font/google";
import { PropsWithChildren } from "react";
import "./globals.css";

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

export default function RootLayout({ children }: PropsWithChildren) {
  return (
    <html lang="en" className="dark">
      <body
        className={cn(font_mono.variable, 'font-display antialiased')}
      >
        <Toaster />
        {children}
      </body>
    </html>
  );
}
