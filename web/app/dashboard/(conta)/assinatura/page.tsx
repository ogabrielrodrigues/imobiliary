import { auth } from "@/actions/auth"
import { ProPlanDialog } from "@/components/pro-plan-dialog"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Separator } from "@/components/ui/separator"
import { Check, Sparkles } from "lucide-react"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Assinatura",
  description: "Assinatura"
}

export default async function SubscriptionPage() {
  const user = await auth()

  return (
    <div>
      <div className="max-w-4xl mx-auto space-y-6">
        <div className="flex justify-between items-center">
          <h1 className="text-3xl font-bold">Assinatura</h1>
        </div>

        <Card>
          <CardHeader>
            <CardDescription>
              Atualmente você está utilizando o Imobiliary no plano <strong className="text-primary">FREE</strong>. Para ter acesso a todas as funcionalidades, assine o plano <strong className="text-primary">PRO</strong>.
            </CardDescription>
          </CardHeader>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Plano {user?.plan.kind.toUpperCase()}</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <div className="flex items-center justify-between">
                <p>{user?.plan.propertiesUsedQuota}/{user?.plan.propertiesTotalQuota}</p>
                <p>Imóveis</p>
              </div>
              <Progress value={(user?.plan.propertiesUsedQuota ?? 0) / (user?.plan.propertiesTotalQuota ?? 1) * 100} />
            </div>

            <Separator />

            <div>
              <ul className="space-y-1">
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Gerencie até 30 imóveis
                </li>
                <li className="flex items-center gap-2 text-sm">
                  <Check className="size-4 text-emerald-500" />
                  Geração de recibos
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
              </ul>
            </div>
          </CardContent>
          <CardFooter className="justify-end">
            <ProPlanDialog>
              <Button>
                <Sparkles />
                Atualizar para o PRO
              </Button>
            </ProPlanDialog>
          </CardFooter>
        </Card>
      </div>
    </div>
  )
}