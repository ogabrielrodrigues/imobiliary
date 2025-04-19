import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Property } from "@/types/property"

interface PropertiesHeaderProps {
  properties: Property[]
}

export function PropertiesHeader({ properties }: PropertiesHeaderProps) {
  const occupied = properties.filter(property => property.status === 'Ocupado')
  const occupied_percent = ((occupied.length / properties.length) * 100) || 0

  const available = properties.filter(property => property.status === 'Disponível')
  const available_percent = ((available.length / properties.length) * 100) || 0

  return (
    <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      <Card className="bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-950 dark:to-emerald-900">
        <CardHeader>
          <CardTitle>Imóveis Cadastrados</CardTitle>
          <CardDescription>Total de imóveis no sistema</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="text-4xl font-bold">{properties.length}</div>
          <p className="text-sm text-muted-foreground mt-2">Total de imóveis no portfólio</p>
        </CardContent>
      </Card>

      <Card className="bg-gradient-to-br from-indigo-50 to-indigo-100 dark:from-indigo-950 dark:to-indigo-900">
        <CardHeader>
          <CardTitle>Imóveis Alugados</CardTitle>
          <CardDescription>Atualmente ocupados</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="text-4xl font-bold">{occupied.length}</div>
          <p className="text-sm text-muted-foreground mt-2">{occupied_percent.toFixed(0)}% do total de imóveis</p>
        </CardContent>
      </Card>

      <Card className="bg-gradient-to-br from-amber-50 to-amber-100 dark:from-amber-950 dark:to-amber-900">
        <CardHeader>
          <CardTitle>Imóveis Disponíveis</CardTitle>
          <CardDescription>Prontos para locação</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="text-4xl font-bold">{available.length}</div>
          <p className="text-sm text-muted-foreground mt-2">{available_percent.toFixed(0)}% do total de imóveis</p>
        </CardContent>
      </Card>
    </div>
  )
}