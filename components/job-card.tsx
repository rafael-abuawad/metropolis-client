import * as React from "react"
import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardHeader,
    CardDescription,
    CardTitle,
    CardContent,
    CardFooter,
} from "@/components/ui/card"
import { Job } from "@/data/job.type"
import { Badge } from "@/components/ui/badge"
import Link from "next/link"

const CardHeaderGradient = React.forwardRef<
    HTMLDivElement,
    React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
    <div
        ref={ref}
        className={cn(
            "h-52 flex flex-col justify-center items-center bg-gradient-to-r from-cyan-500 to-blue-500 rounded-t-xl",
            className,
        )}
        {...props}
    />
))
CardHeaderGradient.displayName = "CardHeaderGradient"

type JobCardProps = React.HTMLAttributes<HTMLDivElement> & Job

const JobCard = ({
    id,
    title,
    company,
    abbr,
    description,
    tags,
    gradient,
    className,
    children,
    color,
    ...props
}: JobCardProps) => {
    return (
        <Link href={`/jobs/${id}`} className="w-full">
            <Card className={cn("w-full", className)} {...props}>
                <CardHeaderGradient className={gradient ? gradient : ""}>
                    {children}
                </CardHeaderGradient>
                <CardHeader>
                    {color === "blue" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-blue-600 dark:text-blue-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "green" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-green-600 dark:text-green-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "ambar" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-ambar-600 dark:text-ambar-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "emerald" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-emerald-600 dark:text-emerald-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "red" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-red-600 dark:text-red-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "pink" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-pink-600 dark:text-pink-500">
                            {company}
                        </CardDescription>
                    )}
                    {color === "violet" && (
                        <CardDescription className="block mb-1 text-xs font-semibold uppercase text-violet-600 dark:text-violet-500">
                            {company}
                        </CardDescription>
                    )}

                    <CardTitle>{title}</CardTitle>
                    <CardDescription>{abbr}</CardDescription>
                </CardHeader>
                <CardContent className="space-y-2">
                    {tags.map((tag, i) => (
                        <Badge key={i} variant="secondary" className="mr-2">
                            {tag.name}
                        </Badge>
                    ))}
                </CardContent>
                <CardContent className="text-gray-500">
                    {description}
                </CardContent>
                <CardFooter className="flex justify-between">
                    <Button variant="outline">Ver Más</Button>
                    <Button>Aplicar</Button>
                </CardFooter>
            </Card>
        </Link>
    )
}

export { JobCard, CardHeaderGradient }
