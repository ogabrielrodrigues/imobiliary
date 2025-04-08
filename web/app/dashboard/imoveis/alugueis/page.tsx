import { Metadata } from "next";

import { SiteHeader } from "@/components/site-header";
import { DataTable } from "@/components/data-table";

import data from './data.json'

export const metadata: Metadata = {
  title: "Alugueis"
}

export default function PropertyRentalsPage() {
  return (
    <>
      <SiteHeader title="Alugueis" />
      <div className="flex flex-1 flex-col">
        <div className="@container/main flex flex-1 flex-col gap-2">
          <div className="h-full flex flex-col gap-4 py-4 md:gap-6 md:py-6">
            <div className="gap-4 px-4 lg:px-6">
              <DataTable data={data} />
            </div>
          </div>
        </div>
      </div>
    </>
  )
}