import { getPlan } from "@/actions/queries/plan/get-plan"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { BarChart, LineChart, PieChart } from "lucide-react"
import { Metadata } from "next"
import { redirect } from "next/navigation"

export const metadata: Metadata = {
  title: "Relatórios",
  description: "Gerencie os relatórios",
}

export default async function RelatoriosPage() {
  const { status, plan } = await getPlan()

  if (status !== 200 || plan?.kind !== "pro") {
    redirect("/auth/login")
  }

  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold mb-6">Relatórios</h1>

      <Tabs defaultValue="financeiro" className="mb-8">
        <TabsList className="mb-4">
          <TabsTrigger value="financeiro" className="flex items-center gap-2">
            <BarChart className="h-4 w-4" />
            <span>Financeiro</span>
          </TabsTrigger>
          <TabsTrigger value="ocupacao" className="flex items-center gap-2">
            <PieChart className="h-4 w-4" />
            <span>Ocupação</span>
          </TabsTrigger>
          <TabsTrigger value="desempenho" className="flex items-center gap-2">
            <LineChart className="h-4 w-4" />
            <span>Desempenho</span>
          </TabsTrigger>
        </TabsList>

        <TabsContent value="financeiro">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Card className="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-950 dark:to-blue-900">
              <CardHeader>
                <CardTitle>Receita Mensal</CardTitle>
                <CardDescription>Aluguéis recebidos</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">R$ 87.450</div>
                <p className="text-sm text-muted-foreground mt-2">Aumento de 5% em relação ao mês anterior</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-green-50 to-green-100 dark:from-green-950 dark:to-green-900">
              <CardHeader>
                <CardTitle>Taxa de Inadimplência</CardTitle>
                <CardDescription>Pagamentos em atraso</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">3,2%</div>
                <p className="text-sm text-muted-foreground mt-2">Redução de 0,8% em relação ao mês anterior</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-orange-50 to-orange-100 dark:from-orange-950 dark:to-orange-900">
              <CardHeader>
                <CardTitle>Repasses aos Proprietários</CardTitle>
                <CardDescription>Valores transferidos</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">R$ 74.332</div>
                <p className="text-sm text-muted-foreground mt-2">85% da receita total</p>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="ocupacao">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Card className="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950 dark:to-purple-900">
              <CardHeader>
                <CardTitle>Taxa de Ocupação</CardTitle>
                <CardDescription>Imóveis alugados</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">92%</div>
                <p className="text-sm text-muted-foreground mt-2">53 de 57 imóveis ocupados</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-pink-50 to-pink-100 dark:from-pink-950 dark:to-pink-900">
              <CardHeader>
                <CardTitle>Tempo Médio de Vacância</CardTitle>
                <CardDescription>Entre locações</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">24 dias</div>
                <p className="text-sm text-muted-foreground mt-2">Redução de 3 dias em relação ao trimestre anterior</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-yellow-50 to-yellow-100 dark:from-yellow-950 dark:to-yellow-900">
              <CardHeader>
                <CardTitle>Renovações de Contrato</CardTitle>
                <CardDescription>Último trimestre</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">87%</div>
                <p className="text-sm text-muted-foreground mt-2">13 de 15 contratos renovados</p>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="desempenho">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Card className="bg-gradient-to-br from-cyan-50 to-cyan-100 dark:from-cyan-950 dark:to-cyan-900">
              <CardHeader>
                <CardTitle>Crescimento Anual</CardTitle>
                <CardDescription>Novos contratos</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">12%</div>
                <p className="text-sm text-muted-foreground mt-2">7 novos imóveis no portfólio</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-indigo-50 to-indigo-100 dark:from-indigo-950 dark:to-indigo-900">
              <CardHeader>
                <CardTitle>Satisfação dos Inquilinos</CardTitle>
                <CardDescription>Pesquisa mensal</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">4,7/5</div>
                <p className="text-sm text-muted-foreground mt-2">Aumento de 0,2 pontos em relação ao mês anterior</p>
              </CardContent>
            </Card>

            <Card className="bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-950 dark:to-emerald-900">
              <CardHeader>
                <CardTitle>Tempo de Resposta</CardTitle>
                <CardDescription>Atendimento a chamados</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-bold">6h</div>
                <p className="text-sm text-muted-foreground mt-2">Redução de 2h em relação ao mês anterior</p>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Relatórios Detalhados</CardTitle>
            <CardDescription>Análises completas e exportáveis</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Relatórios. Aqui você encontrará análises detalhadas sobre o desempenho financeiro, ocupação de imóveis e indicadores gerais.</p>
              <p className="mt-2">Utilize as abas acima para navegar entre diferentes categorias de relatórios e exportar dados para análises mais aprofundadas.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
