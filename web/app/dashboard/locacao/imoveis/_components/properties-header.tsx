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
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
        <div className="absolute z-10 w-1/3 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial from-indigo-500 to-indigo-700" />
        <CardHeader>
          <CardTitle>Imóveis Cadastrados</CardTitle>
          <CardDescription>Total de imóveis no sistema</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="text-4xl font-bold">{properties.length}</div>
          <p className="text-sm text-muted-foreground mt-2">Total de imóveis no portfólio</p>
        </CardContent>
      </Card>

      <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
        <div className="absolute z-10 w-1/3 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial from-amber-500 to-amber-700" />
        <CardHeader>
          <CardTitle>Imóveis Alugados</CardTitle>
          <CardDescription>Atualmente ocupados</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="text-4xl font-bold">{occupied.length}</div>
          <p className="text-sm text-muted-foreground mt-2">{occupied_percent.toFixed(0)}% do total de imóveis</p>
        </CardContent>
      </Card>

      <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden md:col-span-2 xl:col-span-1">
        <div className="absolute z-10 w-1/3 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial from-emerald-500 to-emerald-700" />
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