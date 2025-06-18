import { Toaster } from "@/components/ui/sonner";
import { cn } from "@/lib/utils";
import type { Metadata } from "next";
import { Poppins, Rubik } from "next/font/google";
import { PropsWithChildren } from "react";
import "./globals.css";

const fontSans = Rubik({
  variable: "--font-sans",
  subsets: ["latin"],
});

const fontHeading = Poppins({
  variable: "--font-heading",
  subsets: ["latin"],
  weight: ["700", "800", "900"]
});

export const metadata: Metadata = {
  title: {
    template: "Imobiliary | %s",
    default: 'Imobiliary',
  }
};

export default function RootLayout({ children }: PropsWithChildren) {
  return (
    <html lang="en" className="">
      <body
        className={cn(fontSans.variable, fontHeading.variable, 'font-sans antialiased')}
      >
        <Toaster />
        {children}
      </body>
    </html>
  );
}
