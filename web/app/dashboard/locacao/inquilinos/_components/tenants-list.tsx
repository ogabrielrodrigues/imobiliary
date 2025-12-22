import { Tenant } from "@/types/tenant";
import { TenantCard } from "./tenants-card";

type TenantsListProps = {
  tenants: Tenant[]
}

export function TenantsList({ tenants }: TenantsListProps) {
  return (
    <div className="grid gap-6 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
      {tenants.length > 0
        ? tenants.map((tenant) => <TenantCard key={tenant.id} tenant={tenant} />)
        : (
          <div className="col-span-full text-center py-8">
            <p className="text-muted-foreground">Nenhum inquilino encontrado com os filtros selecionados</p>
          </div>
        )}
    </div>
  )
}