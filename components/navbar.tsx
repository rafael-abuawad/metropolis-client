"use client"

import * as React from "react"

import { Input } from "@/components/ui/input"
import { UserNav } from "@/components/user-nav"
import Link from "next/link"
import { ModeToggle } from "./mode-toggle"
import { Hexagon } from "lucide-react"

interface NavbarProps extends React.HTMLAttributes<HTMLDivElement> {}

export function Navbar({ className, ...props }: NavbarProps) {
    return (
        <nav className="flex-col md:flex">
            <div className="border-b">
                <div className="flex h-16 items-center px-4">
                    <div className="flex items-center space-x-4 lg:space-x-6 md:mx-6">
                        <Link
                            href="/"
                            className="text-sm font-medium transition-colors hover:text-primary"
                        >
                            <span className="flex items-center">
                                <Hexagon className="mr-2 h-6 w-6" />
                                Metropolis
                            </span>
                        </Link>
                        <Link
                            href="/"
                            className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
                        >
                            Empresas
                        </Link>
                    </div>
                    <div className="ml-auto flex items-center space-x-4">
                        <div>
                            <Input
                                type="search"
                                placeholder="Buscar..."
                                className="hidden md:block md:w-[100px] lg:w-[300px]"
                            />
                        </div>
                        <UserNav />
                        <ModeToggle variant="outline" />
                    </div>
                </div>
            </div>
        </nav>
    )
}
