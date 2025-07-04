import { getOwner } from "@/actions/queries/owner/get-owner"
import { getProperty } from "@/actions/queries/property/get-property"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"
import { ArrowLeft, ArrowUpRight } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { notFound } from "next/navigation"
import { Fragment } from "react"
import { bgColorStatusDetail, StatusBadge } from "../_components/status-badge"

type PropertyDetailsPageParams = {
  params: Promise<{ property_id: string }>
}

export const metadata: Metadata = {
  title: "Detalhes do Imóvel",
}

export default async function PropertyDetailsPage({ params }: PropertyDetailsPageParams) {
  const { property_id } = await params
  const { status: status, property } = await getProperty(property_id)

  if (!property || status !== 200) notFound()

  const { status: status_owner, owner } = await getOwner(property.owner_id)

  if (status_owner !== 200) notFound()

  return (
    <Fragment>
      <div className={cn(["absolute z-10 w-1/2 h-3 blur-[92px] top-0 left-1/2 -translate-x-1/2", bgColorStatusDetail(property.status)])} />
      <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
        <div className="flex justify-between gap-4">
          <h1 className="text-2xl font-bold">{property.address.address}</h1>
          <Link href="/dashboard/locacao/imoveis">
            <Button variant="outline">
              <ArrowLeft className="w-4 h-4" />
              <p className="hidden md:block">Voltar</p>
            </Button>
          </Link>
        </div>
        <div className="flex flex-col sm:flex-row gap-2 sm:gap-6 text-xs text-muted-foreground">
          <span>Cód. Água - {property.water_id}</span>
          <span>Cód. Energia - {property.energy_id}</span>
        </div>
        <div className="flex space-x-2">
          <StatusBadge status={property.status} />
          <Badge className="h-8" variant="outline">{property.kind}</Badge>
        </div>

        <Card>
          <CardHeader className="!gap-0 flex justify-between items-center">
            <div className="flex flex-col">
              <CardTitle className="text-lg font-bold">{owner?.fullname}</CardTitle>
              <CardDescription>Proprietário(a)</CardDescription>
            </div>
            <Link href={`/dashboard/locacao/proprietarios/${owner?.id}`}>
              <Button variant="link">
                <ArrowUpRight className="size-4 lg:mr-1" />
                <p className="hidden md:block">Ver detalhes</p>
              </Button>
            </Link>
          </CardHeader>
        </Card>
      </div>
    </Fragment>
  )
}

