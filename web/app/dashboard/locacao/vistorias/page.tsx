import { auth } from "@/actions/auth"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"
import { redirect } from "next/navigation"

export const metadata: Metadata = {
  title: "Vistorias",
  description: "Gerencie as vistorias dos imóveis",
}

export default async function VistoriasPage() {
  const user = await auth()

  if (user?.plan.kind !== 'pro') {
    return redirect('/dashboard')
  }

  return (
    <div className="container py-8">
      <h1 className="text-3xl font-bold mb-6">Vistorias</h1>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Card className="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950 dark:to-purple-900">
          <CardHeader>
            <CardTitle>Total de Vistorias</CardTitle>
            <CardDescription>Vistorias realizadas</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">124</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 8% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-green-50 to-green-100 dark:from-green-950 dark:to-green-900">
          <CardHeader>
            <CardTitle>Vistorias Pendentes</CardTitle>
            <CardDescription>A serem realizadas</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">17</div>
            <p className="text-sm text-muted-foreground mt-2">5 agendadas para esta semana</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-950 dark:to-blue-900">
          <CardHeader>
            <CardTitle>Tempo Médio</CardTitle>
            <CardDescription>Duração das vistorias</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">45 min</div>
            <p className="text-sm text-muted-foreground mt-2">Redução de 10 minutos em relação ao mês anterior</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Gerenciamento de Vistorias</CardTitle>
            <CardDescription>Agende e gerencie todas as vistorias</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Vistorias. Aqui você poderá gerenciar todas as vistorias de entrada e saída dos imóveis.</p>
              <p className="mt-2">Utilize esta seção para agendar novas vistorias, registrar ocorrências e gerar relatórios detalhados.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
