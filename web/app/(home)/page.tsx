import type { Metadata } from "next"
import Link from "next/link"

import { Button } from "@/components/ui/button"
import { Fragment } from "react"
import { HomeHeader } from "./_components/home-header"

export const metadata: Metadata = {
  title: "Home"
}

export default async function HomePage() {
  return (
    <Fragment>
      <HomeHeader />
      <main className="w-screen h-[calc(100svh-132px)] flex flex-col items-center justify-center text-center">
        <div className="px-4 sm:px-0 flex flex-col items-center">
          <h1 className="text-[1.80rem] sm:px-8 leading-8 sm:leading-16 font-extrabold font-heading text-muted mb-4 sm:text-6xl text-left sm:text-center">
            Bem-vindo ao Imobiliary!
          </h1>
          <p className="mb-8 md:text-lg text-base max-w-[640px] text-muted-foreground text-left sm:text-center">
            Aqui você poderá gerenciar seus imóveis com agilidade e eficiência, sem gastar <span className="text-primary font-bold">nenhum real</span> com isso.
          </p>
          <Link href="/login">
            <Button size="lg">Conhecer agora</Button>
          </Link>
        </div>
      </main>
      <footer className=" border-t border-border p-4 h-16 sm:gap-2 w-full flex items-center justify-center text-muted-foreground text-xs sm:text-sm">
        <p>Quer contribuir, e ajudar a aprimorar a plataforma?</p>
        <Link
          target="_blank"
          href="https://github.com/ogabrielrodrigues/imobiliary"
          className="text-primary"
        >
          <Button variant="ghost" className="font-bold">Github</Button>
        </Link>
      </footer>
    </Fragment>
  )
}

