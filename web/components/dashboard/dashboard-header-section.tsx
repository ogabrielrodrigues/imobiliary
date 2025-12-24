import {
	HeaderCard,
	HeaderCardContent,
	HeaderCardHead,
} from "@/components/header-card"
import type { Property } from "@/types/property"

type DashboardHeaderSectionProps = {
	properties: Property[]
}

export function DashboardHeaderSection({
	properties,
}: DashboardHeaderSectionProps) {
	const propertiesCount = properties.length
	const occupied = properties.filter(
		property => property.status === "Ocupado",
	).length
	const occupiedPercent = (occupied / propertiesCount) * 100 || 0

	return (
		<div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
			<HeaderCard>
				<HeaderCardHead
					description="Total de imóveis em locação"
					title="Imóveis Ativos"
				/>
				<HeaderCardContent
					className="text-muted"
					count={propertiesCount.toString().padStart(2, "0")}
				/>
			</HeaderCard>
			<HeaderCard>
				<HeaderCardHead
					description="Total de contratos em vigor"
					title="Contratos Ativos"
				/>
				<HeaderCardContent
					className="text-muted"
					count={occupied.toString().padStart(2, "0")}
				/>
			</HeaderCard>
			<HeaderCard className="md:col-span-2 xl:col-span-1">
				<HeaderCardHead
					description="Percentual de imóveis ocupados"
					title="Taxa de Ocupação"
				/>
				<HeaderCardContent
					className="text-muted"
					count={`${occupiedPercent}%`}
				/>
			</HeaderCard>
		</div>
	)
}
