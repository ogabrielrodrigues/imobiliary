import type { Metadata } from "next"
import Link from "next/link"

import { Button } from "@/components/ui/button"
import { HeroSection } from "./_components/hero-section"
import { HomeHeader } from "./_components/home-header"
import { PlanSection } from "./_components/plan-section"

export const metadata: Metadata = {
  title: "Home"
}

export default async function HomePage() {
  return (
    <div className="w-full overflow-x-hidden">
      <HomeHeader />

      <HeroSection />
      <PlanSection />

      <section className="my-16 flex flex-col items-center">
        <Link href="/login">
          <Button>Conhecer agora</Button>
        </Link>
      </section>
    </div>
  )
}

