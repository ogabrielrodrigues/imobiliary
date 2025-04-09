import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "Contratos",
  description: "Gerencie os contratos de aluguel",
}

export default function ContratosPage() {
  return (
    <div className="container mx-auto py-6">
      <h1 className="text-3xl font-bold tracking-tight">
        <div className="flex flex-col gap-1">
          <span>Contratos</span>
          <span className="text-sm font-normal text-muted-foreground">Gerencie os contratos de aluguel</span>
        </div>
      </h1>

      <div className="grid gap-6 mt-8">
        <Card>
          <CardHeader>
            <CardTitle>Lista de Contratos</CardTitle>
            <CardDescription>Visualize e gerencie todos os contratos ativos e inativos</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="p-8 text-center text-muted-foreground">
              <p>Esta é a página de Contratos. Aqui você poderá gerenciar todos os contratos de aluguel.</p>
              <p className="mt-2">Conteúdo específico da página de Contratos será implementado em breve.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
