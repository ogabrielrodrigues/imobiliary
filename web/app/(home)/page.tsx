import { IconArrowUpRight } from "@tabler/icons-react"
import type { Metadata } from "next"
import Link from "next/link"
import { HomeHeader } from "@/components/home/home-header"
import { Button } from "@/components/ui/button"

export const metadata: Metadata = {
	title: "Home",
}

export default async function HomePage() {
	return (
		<div className="h-screen w-screen flex flex-col">
			<HomeHeader />

			<main className="w-full flex-1 px-4 flex flex-col items-center justify-center text-center">
				<h1 className="text-[1.80rem] sm:px-8 leading-8 sm:leading-16 font-extrabold font-heading text-muted mb-4 sm:text-6xl text-left sm:text-center">
					Bem-vindo ao imobiliary!
				</h1>

				<p className="mb-8 md:text-lg text-base max-w-[640px] text-muted-foreground text-left sm:text-center">
					Aqui você poderá gerenciar seus imóveis com agilidade e eficiência,
					sem gastar <span className="text-primary font-bold">nenhum real</span>{" "}
					com isso!
				</p>

				<Link href="/login">
					<Button>
						<IconArrowUpRight />
						Conhecer agora
					</Button>
				</Link>
			</main>

			<footer className="w-full p-4 gap-1 flex items-center justify-center border-t border-border text-xs sm:text-sm">
				<p className="text-muted-foreground">
					Quer contribuir, e ajudar a aprimorar a plataforma?
				</p>
				<Link
					className="text-primary underline underline-offset-2"
					href="https://github.com/ogabrielrodrigues/imobiliary"
					target="_blank"
				>
					Github
				</Link>
			</footer>
		</div>
	)
}
