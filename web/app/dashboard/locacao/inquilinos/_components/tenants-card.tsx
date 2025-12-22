import { Button } from "@/components/ui/button"
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Tenant } from "@/types/tenant"
import { ArrowUpRight } from "lucide-react"
import Link from "next/link"

type TenantCardProps = {
  tenant: Tenant
}

export function TenantCard({ tenant }: TenantCardProps) {
  return (
    <Card>
      <CardHeader className="flex flex-col h-20">
        <CardTitle>{tenant.fullname}</CardTitle>
        <CardDescription>{tenant.address.address}</CardDescription>
      </CardHeader>
      <CardFooter className="flex justify-end !m-0">
        <div className="flex justify-end">
          <Link href={`/dashboard/locacao/inquilinos/${tenant.id}`}>
            <Button variant="link">
              <ArrowUpRight className="size-4" />Detalhes
            </Button>
          </Link>
        </div>
      </CardFooter>
    </Card>
  )
}