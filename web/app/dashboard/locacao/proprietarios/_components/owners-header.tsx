import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Owner } from "@/types/owner"

interface OwnersHeaderProps {
  owners: Owner[]
  propertiesCount: number
}

export function OwnersHeader({ owners, propertiesCount }: OwnersHeaderProps) {
  const ownersCount = owners.length
  const propertiesPerOwner = propertiesCount / ownersCount
  const propertiesPerOwnerLabel = propertiesCount === 0 ? "Nenhum imóvel cadastrado" : `${propertiesCount} imóveis para ${ownersCount} proprietários`

  return (
    <div className="grid gap-6 md:grid-cols-2">
      <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
        <div className="absolute z-10 w-1/3 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial dark:from-purple-950 dark:to-purple-900" />
        <CardHeader>
          <CardTitle>Total de Proprietários</CardTitle>
          <CardDescription>Proprietários cadastrados</CardDescription>
        </CardHeader>
        <CardContent>
          <strong className="text-4xl font-bold">{ownersCount}</strong>
          <p className="text-sm text-muted-foreground mt-2">Total de proprietários ativos no sistema</p>
        </CardContent>
      </Card>

      <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
        <div className="absolute z-10 w-1/3 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial dark:from-orange-950 dark:to-orange-900" />
        <CardHeader>
          <CardTitle>Imóveis por Proprietário</CardTitle>
          <CardDescription>Média de imóveis</CardDescription>
        </CardHeader>
        <CardContent>
          <strong className="text-4xl font-bold">~{propertiesPerOwner}</strong>
          <p className="text-sm text-muted-foreground mt-2">{propertiesPerOwnerLabel}</p>
        </CardContent>
      </Card>
    </div>
  )
}