import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Alugueres",
  description: "Visão geral dos imóveis alugados",
}

export default function AlugueresPage() {
  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold tracking-tight">
        <div className="flex flex-col gap-1">
          <span>Alugueres</span>
          <span className="text-sm font-normal text-muted-foreground">Visão geral dos imóveis alugados</span>
        </div>
      </h1>

      <div className="grid gap-6 mt-8 md:grid-cols-2 lg:grid-cols-3">
        <Card className="bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-950 dark:to-blue-900">
          <CardHeader>
            <CardTitle>Contratos Ativos</CardTitle>
            <CardDescription>Total de contratos em vigor</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">42</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 8% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-green-50 to-green-100 dark:from-green-950 dark:to-green-900">
          <CardHeader>
            <CardTitle>Receita Mensal</CardTitle>
            <CardDescription>Valor total de aluguéis</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">R$ 68.450</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 5% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-amber-50 to-amber-100 dark:from-amber-950 dark:to-amber-900">
          <CardHeader>
            <CardTitle>Imóveis Disponíveis</CardTitle>
            <CardDescription>Prontos para locação</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">15</div>
            <p className="text-sm text-muted-foreground mt-2">3 novos imóveis este mês</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Página de Aluguéis</CardTitle>
            <CardDescription>Esta é a página principal do módulo de Aluguéis</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Bem-vindo à página principal de Aluguéis. Aqui você encontrará uma visão geral de todos os aluguéis.</p>
              <p className="mt-2">Utilize o menu lateral para navegar entre as diferentes seções do módulo de Aluguéis.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
