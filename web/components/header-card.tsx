import { cn } from "@/lib/utils";
import { PropsWithChildren } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "./ui/card";

type HeaderCardProps = PropsWithChildren & {
  className?: string
}

export function HeaderCard({ className, children }: HeaderCardProps) {
  return (
    <Card className={className}>
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
  className?: string
}

export function HeaderCardContent({ count, className }: HeaderCardContentProps) {
  return (
    <CardContent>
      <div className={cn(["text-4xl font-bold", className])}>{count}</div>
    </CardContent>
  )
}