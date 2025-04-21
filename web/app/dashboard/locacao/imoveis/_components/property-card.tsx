import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"
import { Property } from "@/types/property"
import { ArrowUpRight } from "lucide-react"
import Link from "next/link"
import { bgColorStatusDetail, StatusBadge } from "./status-badge"

type PropertyCardProps = {
  property: Property
}

export function PropertyCard({ property }: PropertyCardProps) {
  return (
    <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
      <div className={cn(["absolute z-10 w-10 h-10 blur-3xl bottom-0 right-0", bgColorStatusDetail(property.status)])} />
      <CardHeader className="flex flex-col gap-4 h-20">
        <CardTitle className="max-h-12">{property.address.mini_address}</CardTitle>
        <div className="flex flex-col gap-2 lg:flex-row lg:items-center">
          <StatusBadge status={property.status} />
          <Badge className="h-8" variant="outline">{property.kind}</Badge>
        </div>
      </CardHeader>
      <CardFooter className="flex justify-end">
        <div className="flex justify-end">
          <Link href={`/dashboard/locacao/imoveis/${property.id}`}>
            <Button variant="outline">
              <ArrowUpRight className="size-4 mr-1" />
              <p className="hidden sm:block">Detalhes</p>
            </Button>
          </Link>
        </div>
      </CardFooter>
    </Card>
  )
}