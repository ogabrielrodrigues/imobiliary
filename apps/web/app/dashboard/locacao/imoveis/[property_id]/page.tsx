import { getOwner } from "@/actions/queries/owner/get-owner"
import { listOwners } from "@/actions/queries/owner/list-owners"
import { getProperty } from "@/actions/queries/property/get-property"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
import { ArrowLeft, ArrowUpRight, EllipsisVertical, Pencil } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { notFound } from "next/navigation"
import { Fragment } from "react"
import { bgColorStatusDetail, StatusBadge } from "../_components/status-badge"
import { AssignPropertyForm } from "./_components/assign-property-form"

type PropertyDetailsPageParams = {
  params: Promise<{ property_id: string }>
}

export const metadata: Metadata = {
  title: "Detalhes do Imóvel",
}

export default async function PropertyDetailsPage({ params }: PropertyDetailsPageParams) {
  const { property_id } = await params
  const { property, status: status } = await getProperty(property_id)

  if (!property || status !== 200) notFound()

  if (property.owner_id !== "") {
    const { status: status_owner, owner } = await getOwner(property.owner_id)

    if (!owner || status_owner !== 200) {
      return
    }

    return (
      <Fragment>
        <div className={cn(["absolute z-10 w-1/2 h-3 blur-[92px] top-0 left-1/2 -translate-x-1/2", bgColorStatusDetail(property.status)])} />
        <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
          <div className="flex justify-between gap-4">
            <h1 className="text-2xl font-bold">{property.address.mini_address}</h1>
            <div className="flex flex-col sm:flex-row gap-2">
              <Link href="/dashboard/locacao/imoveis">
                <Button variant="outline">
                  <ArrowLeft className="w-4 h-4" />
                  <p className="hidden md:block">Voltar</p>
                </Button>
              </Link>
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

          <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
            <CardHeader className="flex justify-between items-center">
              <CardTitle className="text-lg font-bold">Dados do Proprietário</CardTitle>
              <Link href={`/dashboard/locacao/proprietarios/${owner.id}`}>
                <Button variant="outline">
                  <ArrowUpRight className="size-4" />
                  <p className="hidden md:block">Ver detalhes</p>
                </Button>
              </Link>
            </CardHeader>
            <CardContent>
              <div className="flex flex-col gap-2">
                <p className="text-muted-foreground"><strong className="text-primary">Nome:</strong> {owner.fullname}</p>
                <p className="text-muted-foreground"><strong className="text-primary">Email:</strong> {owner.email}</p>
                <p className="text-muted-foreground"><strong className="text-primary">Telefone:</strong> {owner.cellphone}</p>
                <p className="text-muted-foreground"><strong className="text-primary">Endereço:</strong> {owner.address.mini_address}</p>
              </div>
            </CardContent>
          </Card>
        </div>
      </Fragment>
    )
  } else {
    return (
      <Fragment>
        <div className={cn(["absolute z-10 w-1/2 h-3 blur-[92px] top-0 left-1/2 -translate-x-1/2", bgColorStatusDetail(property.status)])} />
        <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
          <div className="flex justify-between gap-4">
            <h1 className="text-2xl font-bold">{property.address.mini_address}</h1>
            <div className="flex flex-col sm:flex-row gap-2">
              <Link href="/dashboard/locacao/imoveis">
                <Button variant="outline">
                  <ArrowLeft className="w-4 h-4" />
                  <p className="hidden md:block">Voltar</p>
                </Button>
              </Link>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button>
                    <EllipsisVertical className="size-4" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent>
                  <DropdownMenuLabel>Opções</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>
                    <Pencil className="size-4 text-white" />Editar
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
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

          <Card>
            <CardContent className="flex flex-col items-center justify-center space-y-6">
              <h1>Esse imóvel ainda não possui um proprietário associado</h1>
              <AssignPropertyForm property_id={property_id} owners={((await listOwners()).owners)} />
            </CardContent>
          </Card>
        </div>
      </Fragment>
    )
  }
}
