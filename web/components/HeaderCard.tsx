import { cn } from "@/lib/utils";
import { PropsWithChildren } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "./ui/card";

type HeaderCardProps = PropsWithChildren & {
  containerClassName?: string
  className?: string
}

export function HeaderCard({ className, containerClassName, children }: HeaderCardProps) {
  return (
    <Card className={cn(["bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden", containerClassName])}>
      <div className={cn(['absolute z-10 w-3/4 h-1/3 blur-[144px] top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2 bg-radial', className])} />
      {children}
    </Card>
  )
}

export type HeaderCardHeadProps = {
  title: string
  description: string
}

export function HeaderCardHead({ title, description }: HeaderCardHeadProps) {
  return (
    <CardHeader>
      <CardTitle>{title}</CardTitle>
      <CardDescription>{description}</CardDescription>
    </CardHeader>
  )
}

type HeaderCardContentProps = {
  count: string
  label: string
  className?: string
}

export function HeaderCardContent({ count, label, className }: HeaderCardContentProps) {
  return (
    <CardContent>
      <div className={cn(["text-4xl font-bold", className])}>{count}</div>
      <p className="text-sm text-muted-foreground mt-2">{label}</p>
    </CardContent>
  )
}