import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Property } from "@/types/property"
import { ArrowUpRight } from "lucide-react"
import Link from "next/link"
import { StatusBadge } from "./status-badge"

type PropertyCardProps = {
  property: Property
}

export function PropertyCard({ property }: PropertyCardProps) {
  return (
    <Card>
      <CardHeader className="flex flex-col gap-4 h-20">
        <CardTitle>{property.address.address}</CardTitle>
        <div className="flex flex-col gap-2 lg:flex-row lg:items-center">
          <StatusBadge status={property.status} />
          <Badge className="h-8" variant="outline">{property.kind}</Badge>
        </div>
      </CardHeader>
      <CardFooter className="flex justify-end">
        <div className="flex justify-end">
          <Link href={`/dashboard/locacao/imoveis/${property.id}`}>
            <Button variant="link">
              <ArrowUpRight className="size-4" />Detalhes
            </Button>
          </Link>
        </div>
      </CardFooter>
    </Card>
  )
}