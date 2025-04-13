import { Card, CardContent, CardFooter } from "@/components/ui/card"

export async function HeroSection() {
  return (
    <section className="mt-16 flex flex-col items-center px-4 pt-10 pb-16 relative">
      <div className="absolute z-10 bg-zinc-50 w-20 h-20 blur-[96px] translate-y-1/2 -translate-x-1/2 top-1/2 left-1/2" />

      <h1 className="text-3xl font-bold sm:text-4xl sm:w-lg md:text-5xl md:w-2xl lg:text-7xl lg:w-[936px] xl:text-8xl xl:w-5xl uppercase select-none">Gerencie sua imobiliária com o Imobiliary!</h1>

      <div className="mt-16 grid grid-cols-1 lg:grid-cols-2 gap-4 sm:w-lg md:w-2xl lg:w-[936px] xl:w-5xl">
        <Card className="bg-zinc-950/40 backdrop-blur-2xl z-20">
          <CardContent>
            <h1 className="text-5xl font-bold md:text-8xl select-none">40%</h1>
          </CardContent>
          <CardFooter>
            <span className="text-xl md:text-3xl text-muted-foreground select-none">de economia de tempo na realização de tarefas repetitivas.</span>
          </CardFooter>
        </Card>
        <Card className="bg-zinc-950/40 backdrop-blur-2xl">
          <CardContent>
            <h1 className="text-5xl font-bold md:text-8xl select-none">Menor</h1>
          </CardContent>
          <CardFooter>
            <span className="text-xl md:text-3xl text-muted-foreground select-none">custo com administração e outras soluções de gerenciamento.</span>
          </CardFooter>
        </Card>

        <Card className="bg-zinc-950/40 backdrop-blur-2xl">
          <CardContent>
            <h1 className="text-5xl font-bold md:text-8xl select-none">Maior</h1>
          </CardContent>
          <CardFooter>
            <span className="text-xl md:text-3xl text-muted-foreground select-none">eficiência e performance no gerenciamento de imóveis.</span>
          </CardFooter>
        </Card>
        <Card className="bg-zinc-950/40 backdrop-blur-2xl z-20">
          <CardContent>
            <h1 className="text-5xl font-bold md:text-8xl select-none">Simples</h1>
          </CardContent>
          <CardFooter>
            <span className="text-xl md:text-3xl text-muted-foreground select-none">e intuitivo para o gerenciamento de sua imobiliária.</span>
          </CardFooter>
        </Card>
      </div>
    </section>
  )
}
