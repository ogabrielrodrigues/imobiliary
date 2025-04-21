import { Owner } from "@/types/owner";
import { OwnerCard } from "./owner-card";

type OwnersListProps = {
  owners: Owner[]
}

export function OwnersList({ owners }: OwnersListProps) {
  return (
    <div className="grid gap-6 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
      {owners.length > 0
        ? owners.map((owner) => <OwnerCard key={owner.id} owner={owner} />)
        : (
          <div className="col-span-full text-center py-8">
            <p className="text-muted-foreground">Nenhum propretário encontrado com os busca desejada.</p>
          </div>
        )}
    </div>
  )
}