import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/HeaderCard"

type TenantsHeaderProps = {}

export function TenantsHeader({ }: TenantsHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard className="from-purple-500 to-purple-700">
        <HeaderCardHead
          title="Total de Inquilinos"
          description="Inquilinos ativos no sistema"
        />
        <HeaderCardContent
          count="78"
          label="Aumento de 12% em relação ao mês anterior"
          className="text-purple-100"
        />
      </HeaderCard>
      <HeaderCard className="from-pink-500 to-pink-700">
        <HeaderCardHead
          title="Novos Inquilinos"
          description="Adicionados este mês"
        />
        <HeaderCardContent
          count="9"
          label="3 mais que no mês anterior"
          className="text-pink-100"
        />
      </HeaderCard>
      <HeaderCard className="from-sky-500 to-sky-700" containerClassName="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Inquilinos Pontuais"
          description="Pagamentos em dia"
        />
        <HeaderCardContent
          count="92%"
          label="Aumento de 5% em relação ao mês anterior"
          className="text-sky-100"
        />
      </HeaderCard>
    </div>
  )
}