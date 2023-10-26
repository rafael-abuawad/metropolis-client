import * as React from "react"

import { cn } from "@/lib/utils"

const Hero = React.forwardRef<
    HTMLDivElement,
    React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
    return (
        <div className="relative overflow-hidden">
            <div className="max-w-[85rem] mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-24">
                <div
                    ref={ref}
                    className={cn("text-center", className)}
                    {...props}
                />
            </div>
        </div>
    )
})
Hero.displayName = "Hero"

const HeroTitle = React.forwardRef<
    HTMLHeadingElement,
    React.HTMLAttributes<HTMLHeadingElement>
>(({ className, ...props }, ref) => {
    return (
        <h1
            ref={ref}
            className={cn(
                "text-4xl sm:text-6xl font-bold text-gray-800 dark:text-gray-200",
                className,
            )}
            {...props}
        />
    )
})
HeroTitle.displayName = "HeroTitle"

const HeroSubtitle = React.forwardRef<
    HTMLParagraphElement,
    React.HTMLAttributes<HTMLParagraphElement>
>(({ className, ...props }, ref) => {
    return (
        <h1
            ref={ref}
            className={cn("mt-3 text-gray-600 dark:text-gray-400", className)}
            {...props}
        />
    )
})
HeroSubtitle.displayName = "HeroSubtitle"

const HeroContent = React.forwardRef<
    HTMLDivElement,
    React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
    return (
        <div
            ref={ref}
            className={cn("mt-7 sm:mt-12 mx-auto max-w-xl", className)}
            {...props}
        />
    )
})
HeroContent.displayName = "HeroContent"

export { Hero, HeroTitle, HeroSubtitle, HeroContent }
