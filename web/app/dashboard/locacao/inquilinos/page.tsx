import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Inquilinos",
  description: "Gerencie os inquilinos dos imóveis",
}

export default function InquilinosPage() {
  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold tracking-tight">
        <div className="flex flex-col gap-1">
          <span>Inquilinos</span>
          <span className="text-sm font-normal text-muted-foreground">Gerencie os inquilinos dos imóveis</span>
        </div>
      </h1>

      <div className="grid gap-6 mt-8 md:grid-cols-2 lg:grid-cols-3">
        <Card className="bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950 dark:to-purple-900">
          <CardHeader>
            <CardTitle>Total de Inquilinos</CardTitle>
            <CardDescription>Inquilinos ativos no sistema</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">78</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 12% em relação ao mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-pink-50 to-pink-100 dark:from-pink-950 dark:to-pink-900">
          <CardHeader>
            <CardTitle>Novos Inquilinos</CardTitle>
            <CardDescription>Adicionados este mês</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">9</div>
            <p className="text-sm text-muted-foreground mt-2">3 mais que no mês anterior</p>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-br from-cyan-50 to-cyan-100 dark:from-cyan-950 dark:to-cyan-900">
          <CardHeader>
            <CardTitle>Inquilinos Pontuais</CardTitle>
            <CardDescription>Pagamentos em dia</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="text-4xl font-bold">92%</div>
            <p className="text-sm text-muted-foreground mt-2">Aumento de 5% em relação ao mês anterior</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Gerenciamento de Inquilinos</CardTitle>
            <CardDescription>Cadastre e gerencie todos os inquilinos</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Inquilinos. Aqui você poderá gerenciar todos os inquilinos cadastrados no sistema.</p>
              <p className="mt-2">Utilize esta seção para adicionar novos inquilinos, visualizar informações detalhadas e gerenciar contratos associados.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
