import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Proprietários",
  description: "Gerencie os proprietários dos imóveis",
}

export default function ProprietariosPage() {
  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold tracking-tight">
        <div className="flex flex-col gap-1">
          <span>Proprietários</span>
          <span className="text-sm font-normal text-muted-foreground">Gerenciamento de proprietários de imóveis</span>
        </div>
      </h1>

      <div className="grid gap-6 mt-8 md:grid-cols-2 lg:grid-cols-3">
        <Card className="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950 dark:to-purple-900">
          <CardHeader>
            <CardTitle>Total de Proprietários</CardTitle>
            <CardDescription>Proprietários cadastrados</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">35</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 3% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-950 dark:to-emerald-900">
          <CardHeader>
            <CardTitle>Imóveis por Proprietário</CardTitle>
            <CardDescription>Média de imóveis</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">1,6</div>
            <p className="text-sm text-muted-foreground mt-2">57 imóveis para 35 proprietários</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-orange-50 to-orange-100 dark:from-orange-950 dark:to-orange-900">
          <CardHeader>
            <CardTitle>Repasses Pendentes</CardTitle>
            <CardDescription>Pagamentos a realizar</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">R$ 32.450</div>
            <p className="text-sm text-muted-foreground mt-2">Repasses programados para este mês</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Gerenciamento de Proprietários</CardTitle>
            <CardDescription>Cadastre e gerencie todos os proprietários</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Proprietários. Aqui você poderá gerenciar todos os proprietários de imóveis cadastrados no sistema.</p>
              <p className="mt-2">Utilize esta seção para adicionar novos proprietários, visualizar informações detalhadas e gerenciar os imóveis associados.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
