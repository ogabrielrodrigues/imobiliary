import type { Metadata } from "next"
import Link from "next/link"

import { HousePlus } from "lucide-react"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

export const metadata: Metadata = {
  title: "Imobiliary | Home"
}

export default function HomePage() {
  return (
    <div className="w-full overflow-x-hidden">
      <header className="fixed top-0 left-0 z-50 w-full bg-zinc-950/95 backdrop-blur-3xl p-4 flex items-center justify-center gap-4">
        <Link href="/login">
          <Button variant="ghost">login</Button>
        </Link>
        <div className="flex items-center gap-2">
          <HousePlus className="size-8" />
          <span className="sr-only">Imobiliary</span>
          <h1 className="text-lg font-bold hidden sm:block">Imobiliary</h1>
        </div>
        <Link href="/cadastro">
          <Button variant="ghost">cadastro</Button>
        </Link>
      </header>
      <main className="mt-16 flex flex-col items-center px-4 py-10 relative">
        <div className="absolute z-10 bg-emerald-600 w-20 h-20 blur-[96px] translate-y-1/2 -translate-x-1/2 top-1/2 left-1/2" />

        <h1 className="text-3xl font-bold md:text-5xl md:w-3xl">Gerencie sua imobiliária com o Imobiliary!</h1>

        <div className="mt-16 grid grid-cols-1 md:grid-cols-2 gap-4">
          <Card className="bg-zinc-950/40 backdrop-blur-2xl max-w-md z-20">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl md:w-3xl">40%</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground">de economia de tempo na realização de tarefas repetitivas.</span>
            </CardFooter>
          </Card>


          <Card className=" bg-zinc-950/40 backdrop-blur-2xl max-w-md">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl md:w-3xl">- Menos</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground">custo com administração e outras soluções de gerenciamento.</span>
            </CardFooter>
          </Card>
          <Card className="bg-zinc-950/40 backdrop-blur-2xl max-w-md">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl md:w-3xl">+ Mais</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground">eficiência no gerenciamento de imóveis.</span>
            </CardFooter>
          </Card>
          <Card className="bg-zinc-950/40 backdrop-blur-2xl max-w-md z-20">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl md:w-3xl">Simples</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground">Plataforma simples e intuitiva para o gerenciamento de sua imobiliária.</span>
            </CardFooter>
          </Card>
        </div>
      </main>
      <section className="flex flex-col items-center px-4 py-10 relative">
        <h1 className="text-3xl font-bold md:text-5xl ">Entre em contato</h1>
      </section>
    </div>
  )
}
