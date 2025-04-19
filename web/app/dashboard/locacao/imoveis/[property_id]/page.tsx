import { getProperty } from "@/actions/properties"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { ArrowLeft, EllipsisVertical } from "lucide-react"
import Link from "next/link"
import { StatusBadge } from "../_components/status-badge"

type PropertyDetailsPageParams = {
  params: Promise<{ property_id: string }>
}

export default async function PropertyDetailsPage({ params }: PropertyDetailsPageParams) {
  const { property_id } = await params
  const property = await getProperty(property_id)

  if (!property) {
    return <div className="container mx-auto">Imóvel não encontrado</div>
  }

  return <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
    <div className="flex justify-between gap-4">
      <h1 className="text-2xl font-bold">{property.address.mini_address}</h1>
      <div className="flex flex-col sm:flex-row gap-2">
        <Link href="/dashboard/locacao/imoveis">
          <Button variant="outline">
            <ArrowLeft className="w-4 h-4" />
            <p className="hidden md:block">Voltar</p>
          </Button>
        </Link>
        <Button>
          <EllipsisVertical className="size-4" />
        </Button>
      </div>
    </div>
    <div className="flex flex-col sm:flex-row gap-2 sm:gap-6 text-xs text-muted-foreground">
      <span>Cód. Água - {property.water_id}</span>
      <span>Cód. Energia - {property.energy_id}</span>
    </div>
    <div className="flex space-x-2">
      <StatusBadge status={property.status} />
      <Badge className="h-8" variant="outline">{property.kind}</Badge>
    </div>
  </div>
}