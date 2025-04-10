import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Check, Sparkles, X } from "lucide-react"

export function PlanSection() {
  return (
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
  )
}
