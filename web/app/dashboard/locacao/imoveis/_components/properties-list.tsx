import { Property } from "@/types/property";
import { PropertyCard } from "./property-card";

type PropertyListProps = {
  properties: Property[]
}

export function PropertyList({ properties }: PropertyListProps) {
  return (
    <div className="grid gap-6 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
      {properties.length > 0 ? (
        properties.map((property) => (
          <PropertyCard key={property.id} property={property} />
        ))
      ) : (
        <div className="col-span-full text-center py-8">
          <p className="text-muted-foreground">Nenhum imóvel encontrado com os filtros selecionados.</p>
        </div>
      )}
    </div>
  )
}