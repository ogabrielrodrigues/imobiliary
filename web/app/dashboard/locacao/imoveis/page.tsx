import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default function ImoveisPage() {
  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold tracking-tight">
        <div className="flex flex-col gap-1">
          <span>Imóveis</span>
          <span className="text-sm font-normal text-muted-foreground">Gerencie os imóveis disponíveis para aluguel</span>
        </div>
      </h1>

      <div className="grid gap-6 mt-8 md:grid-cols-2 lg:grid-cols-3">
        <Card className="bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-950 dark:to-emerald-900">
          <CardHeader>
            <CardTitle>Imóveis Cadastrados</CardTitle>
            <CardDescription>Total de imóveis no sistema</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">57</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 5% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-indigo-50 to-indigo-100 dark:from-indigo-950 dark:to-indigo-900">
          <CardHeader>
            <CardTitle>Imóveis Alugados</CardTitle>
            <CardDescription>Atualmente ocupados</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">42</div>
            <p className="text-sm text-muted-foreground mt-2">74% do total de imóveis</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-amber-50 to-amber-100 dark:from-amber-950 dark:to-amber-900">
          <CardHeader>
            <CardTitle>Imóveis Disponíveis</CardTitle>
            <CardDescription>Prontos para locação</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">15</div>
            <p className="text-sm text-muted-foreground mt-2">26% do total de imóveis</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Gerenciamento de Imóveis</CardTitle>
            <CardDescription>Cadastre e gerencie todos os imóveis</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Imóveis. Aqui você poderá gerenciar todos os imóveis disponíveis para aluguel.</p>
              <p className="mt-2">Utilize esta seção para adicionar novos imóveis, atualizar informações e verificar a disponibilidade.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
