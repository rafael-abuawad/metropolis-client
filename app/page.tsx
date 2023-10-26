import { Hero, HeroContent, HeroSubtitle, HeroTitle } from "@/components/hero"
import { Icons } from "@/components/icons"
import { JobCard } from "@/components/job-card"
import { Navbar } from "@/components/navbar"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
    CodeIcon,
    BookIcon,
    BrushIcon,
    MegaphoneIcon,
    UserIcon,
    CandlestickChart,
    Martini,
    Stethoscope,
    Wrench,
    HomeIcon,
} from "lucide-react"

export default function Home() {
    const tags = [
        {
            name: "Information Technology",
            icon: <CodeIcon className="mr-2" size={16} />,
        },
        {
            name: "Healthcare and Medicine",
            icon: <Stethoscope className="mr-2" size={16} />,
        },
        {
            name: "Finance and Accounting",
            icon: <CandlestickChart className="mr-2" size={16} />,
        },
        { name: "Education", icon: <BookIcon className="mr-2" size={16} /> },
        { name: "Engineering", icon: <Wrench className="mr-2" size={16} /> },
        {
            name: "Graphic Design",
            icon: <BrushIcon className="mr-2" size={16} />,
        },
        {
            name: "Sales and Marketing",
            icon: <MegaphoneIcon className="mr-2" size={16} />,
        },
        {
            name: "Human Resources",
            icon: <UserIcon className="mr-2" size={16} />,
        },
        {
            name: "Construction and Architecture",
            icon: <HomeIcon className="mr-2" size={16} />,
        },
        { name: "Food Services", icon: <Martini className="mr-2" size={16} /> },
    ]

    return (
        <>
            <Navbar />
            <main className="flex flex-col items-center justify-between p-4 lg:p-24">
                <Hero>
                    <HeroTitle>Metropolis</HeroTitle>
                    <HeroSubtitle>
                        Con nuestros filtros, encuentra cientos de trabajos en
                        diferentes áreas y localidades, que se ajustan a tus
                        preferencias y habilidades.
                    </HeroSubtitle>
                    <HeroContent>
                        <div className="relative">
                            {/* <!-- Form --> */}
                            <form className="z-10">
                                <Input
                                    type="email"
                                    name="hs-search-article-1"
                                    id="hs-search-article-1"
                                    className="p-6 text-lg md:text-2xl block w-full z-10"
                                    placeholder="Search article"
                                />
                            </form>
                            {/* <!-- End Form --> */}

                            {/* <!-- SVG Element --> */}
                            <div className="hidden md:block absolute top-0 right-0 -translate-y-12 translate-x-12">
                                <svg
                                    className="w-16 h-auto text-orange-500"
                                    width="121"
                                    height="135"
                                    viewBox="0 0 121 135"
                                    fill="none"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path
                                        d="M5 16.4754C11.7688 27.4499 21.2452 57.3224 5 89.0164"
                                        stroke="currentColor"
                                        strokeWidth="10"
                                        strokeLinecap="round"
                                    />
                                    <path
                                        d="M33.6761 112.104C44.6984 98.1239 74.2618 57.6776 83.4821 5"
                                        stroke="currentColor"
                                        strokeWidth="10"
                                        strokeLinecap="round"
                                    />
                                    <path
                                        d="M50.5525 130C68.2064 127.495 110.731 117.541 116 78.0874"
                                        stroke="currentColor"
                                        strokeWidth="10"
                                        strokeLinecap="round"
                                    />
                                </svg>
                            </div>
                            {/* <!-- End SVG Element --> */}

                            {/* <!-- SVG Element --> */}
                            <div className="hidden md:block absolute bottom-0 left-0 translate-y-10 -translate-x-36">
                                <svg
                                    className="w-40 h-auto text-cyan-500"
                                    width="347"
                                    height="188"
                                    viewBox="0 0 347 188"
                                    fill="none"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path
                                        d="M4 82.4591C54.7956 92.8751 30.9771 162.782 68.2065 181.385C112.642 203.59 127.943 78.57 122.161 25.5053C120.504 2.2376 93.4028 -8.11128 89.7468 25.5053C85.8633 61.2125 130.186 199.678 180.982 146.248L214.898 107.02C224.322 95.4118 242.9 79.2851 258.6 107.02C274.299 134.754 299.315 125.589 309.861 117.539L343 93.4426"
                                        stroke="currentColor"
                                        strokeWidth="7"
                                        strokeLinecap="round"
                                    />
                                </svg>
                            </div>
                            {/* <!-- End SVG Element --> */}
                        </div>

                        <div className="mt-10 sm:mt-20 space-x-2 space-y-2">
                            {tags.map((tag, i) => (
                                <Button
                                    key={i}
                                    variant="outline"
                                    className="text-sm md:text-base"
                                >
                                    {tag.icon} {tag.name}
                                </Button>
                            ))}
                        </div>
                    </HeroContent>
                </Hero>

                {/* <!-- Card Blog --> */}
                <div className="max-w-[85rem] px-4 py-10 sm:px-6 lg:px-8 lg:py-14 mx-auto">
                    {/* <!-- Grid --> */}
                    <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
                        {/* <!-- Card --> */}
                        <JobCard
                            abbr="Crece en la nube"
                            company="Attlasian"
                            title="Clud Manager"
                            description="Lorem ipsum dolor sit amet consectetur adipisicing elit. Deserunt quis, repellat fugiat, error, magnam recusandae quam ipsum iste minus similique voluptatum dolorem inventore necessitatibus voluptatem iure? Maiores velit sapiente veritatis.m"
                            tags={tags}
                            color="blue"
                        >
                            <Icons.atlassian className="w-28 h-28" />
                        </JobCard>
                        {/* <!-- End Card --> */}

                        <JobCard
                            abbr="Se parte de lo mejor"
                            company="Apple"
                            title="Ingeniero De Sistemas"
                            description="Lorem ipsum dolor sit amet consectetur adipisicing elit. Deserunt quis, repellat fugiat, error, magnam recusandae quam ipsum iste minus similique voluptatum dolorem inventore necessitatibus voluptatem iure? Maiores velit sapiente veritatis.m"
                            tags={tags}
                            gradient="from-indigo-500 via-purple-500 to-pink-500"
                            color="violet"
                        >
                            <Icons.apple className="w-28 h-28 text-white" />
                        </JobCard>
                    </div>
                    {/* <!-- End Grid --> */}
                </div>
                {/* <!-- End Card Blog --> */}
            </main>
        </>
    )
}
