import type { Metadata } from "next"
import { Poppins, Rubik } from "next/font/google"
import type { PropsWithChildren } from "react"
import { Toaster } from "@/components/ui/sonner"
import { cn } from "@/lib/utils"
import "./globals.css"

const fontSans = Rubik({
	variable: "--font-sans",
	subsets: ["latin"],
})

const fontHeading = Poppins({
	variable: "--font-heading",
	subsets: ["latin"],
	weight: ["700", "800", "900"],
})

export const metadata: Metadata = {
	title: {
		template: "imobiliary | %s",
		default: "imobiliary",
	},
}

export default function RootLayout({ children }: PropsWithChildren) {
	return (
		<html lang="pt-BR">
			<body
				className={cn(
					fontSans.variable,
					fontHeading.variable,
					"font-sans antialiased !bg-primary/5 w-svw min-h-svh",
				)}
			>
				<Toaster />
				{children}
			</body>
		</html>
	)
}
