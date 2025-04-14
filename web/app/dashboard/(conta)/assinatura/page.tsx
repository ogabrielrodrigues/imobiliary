import { auth } from "@/actions/auth"
import { FreePlan, ProPlan } from "@/components/plan-describe"
import { ProPlanDialog } from "@/components/pro-plan-dialog"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Separator } from "@/components/ui/separator"
import { Sparkles } from "lucide-react"
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
              {user?.plan.kind === "free"
                ? (<>Atualmente você está utilizando o Imobiliary no plano <strong className="text-primary">FREE</strong>.<br /><br />Para ter acesso a todas as funcionalidades, assine o plano <strong className="text-primary">PRO</strong>.</>)
                : (<>Parabéns! Você já possui o plano <strong className="text-primary">{user?.plan.kind.toUpperCase()}</strong> e já tendo acesso a todas as funcionalidades.</>)
              }
            </CardDescription>
          </CardHeader>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Plano {user?.plan.kind.toUpperCase()}</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            {user?.plan.kind === 'free' && (
              <>
                <div className="space-y-2">
                  <div className="flex items-center justify-between">
                    <p className="inline-flex items-center">{user?.plan.properties_used_quota === 0 ? "0" : user?.plan.properties_used_quota.toString().padStart(2, "0")}/{user?.plan.properties_total_quota.toString().padStart(2, "0")}</p>
                    <p>imóveis</p>
                  </div>
                  <Progress value={(user?.plan.properties_used_quota ?? 0) / (user?.plan.properties_total_quota ?? 1) * 100} />
                </div>

                <Separator />
              </>
            )}

            <div>
              <ul className="space-y-1">
                {user?.plan.kind === "free"
                  ? (<FreePlan />)
                  : (<ProPlan />)}
              </ul>
            </div>
          </CardContent>
          {user?.plan.kind === "free" && (
            <CardFooter className="justify-end">
              <ProPlanDialog>
                <Button>
                  <Sparkles />
                  Atualizar para o PRO
                </Button>
              </ProPlanDialog>
            </CardFooter>
          )}
        </Card>
      </div>
    </div>
  )
}