import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"
import { Tenant } from "@/types/tenant"

type TenantsHeaderProps = {
  tenants: Tenant[]
}

export function TenantsHeader({ tenants }: TenantsHeaderProps) {
  const tenantsCount = tenants.length
  const exactRentPay = 0

  return (
    <div className="grid gap-6 md:grid-cols-2">
      <HeaderCard>
        <HeaderCardHead
          title="Total de Inquilinos"
          description="Inquilinos ativos no sistema"
        />
        <HeaderCardContent
          count={tenantsCount.toString().padStart(2, '0')}
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard className="">
        <HeaderCardHead
          title="Inquilinos Pontuais"
          description="Pagamentos em dia"
        />
        <HeaderCardContent
          count={`${exactRentPay.toString()}%`}
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}