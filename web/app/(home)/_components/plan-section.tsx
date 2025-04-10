import { FreePlan, ProPlan } from "@/components/plan-describe"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Sparkles } from "lucide-react"

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
            <FreePlan />
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
            <ProPlan />
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
