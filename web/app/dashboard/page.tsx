import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import type { Metadata } from "next"

export const metadata: Metadata = {
  title: "Dashboard"
}

export default function DashboardPage() {
  return (
    <div className="container mx-auto">
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        <Card className="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-950 dark:to-blue-900">
          <CardHeader>
            <CardTitle>Imóveis Ativos</CardTitle>
            <CardDescription>Total de imóveis em locação</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">253</div>
            <p className="text-sm text-muted-foreground mt-2">+8.3% desde o mês passado</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950 dark:to-purple-900">
          <CardHeader>
            <CardTitle>Contratos Ativos</CardTitle>
            <CardDescription>Total de contratos em vigor</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">87</div>
            <p className="text-sm text-muted-foreground mt-2">+5.2% desde o mês passado</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-amber-50 to-amber-100 dark:from-amber-950 dark:to-amber-900">
          <CardHeader>
            <CardTitle>Taxa de Ocupação</CardTitle>
            <CardDescription>Percentual de imóveis ocupados</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">85.5%</div>
            <p className="text-sm text-muted-foreground mt-2">+2.3% desde o mês passado</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Desempenho de Locação por Região</CardTitle>
            <CardDescription>Análise do desempenho das locações por região</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="w-full flex items-center justify-center py-16">
              <p className="text-muted-foreground">Gráfico de desempenho de locação por região será exibido aqui</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}