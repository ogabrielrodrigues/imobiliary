import type { Metadata } from "next"
import Link from "next/link"

import { auth, logout } from "@/actions/auth"
import { Avatar, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip"
import { Check, HousePlus, LogOut, Sparkles, X } from "lucide-react"
export const metadata: Metadata = {
  title: "Imobiliary | Home"
}

export default async function HomePage() {
  const user = await auth()

  return (
    <div className="w-full overflow-x-hidden">
      <header className="fixed top-0 left-0 z-50 w-full bg-zinc-950/95 backdrop-blur-3xl p-4 flex items-center justify-between gap-4">
        <div className="absolute right-8">
          {user && <div className="flex items-center space-x-4">
            <Tooltip>
              <TooltipTrigger asChild>
                <Link href="/dashboard/conta">
                  <Avatar>
                    <AvatarImage src={user.avatar} />
                  </Avatar>
                </Link>
              </TooltipTrigger>
              <TooltipContent>
                <p>Ver Conta</p>
              </TooltipContent>
            </Tooltip>
            <Button asChild variant="ghost">
              <Link href="/dashboard">
                Dashboard
              </Link>
            </Button>
            <form action={logout}>
              <Button variant="outline" type="submit">
                <LogOut className="size-4" />
                <p className="hidden sm:block">Sair</p>
              </Button>
            </form>
          </div>}
          {!user && <div>
            <Link href="/login">
              <Button variant="ghost">Login</Button>
            </Link>
            <Link href="/cadastro">
              <Button variant="ghost">Cadastro</Button>
            </Link>
          </div>}
        </div>
        <div className="flex items-center gap-2">
          <HousePlus className="size-5" />
          <span className="sr-only">Imobiliary</span>
          <h1 className="text-lg font-bold hidden sm:block select-none">Imobiliary</h1>
        </div>
      </header>
      <main className="mt-16 flex flex-col items-center px-4 pt-10 pb-16 relative">
        <div className="absolute z-10 bg-zinc-50 w-20 h-20 blur-[96px] translate-y-1/2 -translate-x-1/2 top-1/2 left-1/2" />

        <h1 className="text-3xl font-bold sm:text-4xl sm:w-lg md:text-5xl md:w-2xl lg:text-7xl lg:w-[936px] xl:text-8xl xl:w-5xl uppercase select-none">Gerencie sua imobiliária com o Imobiliary!</h1>

        <div className="mt-16 grid grid-cols-1 lg:grid-cols-2 gap-4 sm:w-lg md:w-2xl lg:w-[936px] xl:w-5xl">
          <Card className="bg-zinc-950/40 backdrop-blur-2xl z-20">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl select-none">40%</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground select-none">de economia de tempo na realização de tarefas repetitivas.</span>
            </CardFooter>
          </Card>
          <Card className="bg-zinc-950/40 backdrop-blur-2xl">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl select-none">Menor</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground select-none">custo com administração e outras soluções de gerenciamento.</span>
            </CardFooter>
          </Card>

          <Card className="bg-zinc-950/40 backdrop-blur-2xl">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl select-none">Maior</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground select-none">eficiência e performance no gerenciamento de imóveis.</span>
            </CardFooter>
          </Card>
          <Card className="bg-zinc-950/40 backdrop-blur-2xl z-20">
            <CardContent>
              <h1 className="text-5xl font-bold md:text-8xl select-none">Simples</h1>
            </CardContent>
            <CardFooter>
              <span className="text-xl md:text-3xl text-muted-foreground select-none">e intuitivo para o gerenciamento de sua imobiliária.</span>
            </CardFooter>
          </Card>
        </div>
      </main>

      <section className="w-full flex flex-col items-center px-4 space-y-4">
        <div className="flex flex-col gap-4">
          <h1 className="text-3xl font-bold">Planos</h1>
        </div>

        <div className="flex flex-col gap-4 md:flex-row">
          <Card className="md:w-1/2">
            <CardHeader className="md:h-[100px]">
              <CardTitle className="text-2xl">Plano FREE</CardTitle>
              <CardDescription className="text-sm">
                Ideal para quem está começando e quer experimentar o Imobiliary.
              </CardDescription>
            </CardHeader>
            <CardContent className="md:h-60">
              <ul className="space-y-1">
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Gerenciamento de até 30 imóveis
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Métricas e relatórios simples
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Até 2 usuários
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Suporte via e-mail
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <X className="size-4 text-red-500" />
                  Notificações por e-mail e push
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <X className="size-4 text-red-500" />
                  Levantamento de taxas de consumo
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <X className="size-4 text-red-500" />
                  Gerencimento de vistorias
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <X className="size-4 text-red-500" />
                  Administração de carteira
                </li>
              </ul>
            </CardContent>
            <CardFooter className="flex flex-row justify-end">
              <h1 className="text-2xl font-bold">Grátis</h1>
            </CardFooter>
          </Card>

          <Card className="md:w-1/2">
            <CardHeader className="md:h-[100px]">
              <CardTitle className="text-2xl flex gap-2">Plano PRO<Sparkles className="size-5" /></CardTitle>
              <CardDescription className="text-sm">
                Ideal para quem já tem uma imobiliária e quer otimizar o gerenciamento.
              </CardDescription>
            </CardHeader>
            <CardContent className="md:h-60">
              <ul className="space-y-1">
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Gerenciamento ilimitado de imóveis
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Notificações por e-mail e push
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Levantamento de taxas de consumo
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Gerencimento de vistorias
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Administração de carteira
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Métricas e relatórios detalhados
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Até 5 usuários
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Suporte 24/7 via e-mail e chat
                </li>
              </ul>
            </CardContent>
            <CardFooter className="flex flex-row justify-end">
              <h1 className="text-2xl font-bold">R$15,99</h1>
              <p className="text-xl text-muted-foreground">/mês</p>
            </CardFooter>
          </Card>
        </div>
      </section>

      <section className="my-16 flex flex-col items-center">
        <Link href="/login">
          <Button>Conhecer agora</Button>
        </Link>
      </section>
    </div>
  )
}

