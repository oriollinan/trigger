import Link from "next/link";
import React from "react";
import { GrDocumentImage } from "react-icons/gr";
import { MdAddBox, MdLogout } from "react-icons/md";
import { SiGooglegemini } from "react-icons/si";
import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react"
import { Button } from "@/components/ui/button";
import { newTrigger } from "@/app/home/lib/new-trigger";

export const SideMenu = ({workspaceLen}: {workspaceLen: number}) => {
  const [loading, setLoading] = React.useState<boolean>(false);
  const links = [
    {
      name: "Templates",
      href: "/templates",
      icon: (
        <GrDocumentImage className="text-neutral-700 dark:text-neutral-200 mr-2 w-5 h-5" />
      ),
    },
    {
      name: "Triggers",
      href: "/home",
      icon: (
        <SiGooglegemini className="text-neutral-700 dark:text-neutral-200 mr-2 w-5 h-5" />
      ),
    },
  ];

  const mutation = useMutation({
    mutationFn: newTrigger,
    onSuccess: (trigger) => {
      window.location.href = `/trigger/${trigger.id}`;
    },
  });

  const handleClick = () => {
    setLoading(true)
    mutation.mutate({name: `Workspace ${workspaceLen}`, nodes: []});
  };

  return (
    <div className="hidden md:flex flex-col h-full p-7">
      <div className="p-4">
        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">
          Dashboard
        </h2>
      </div>
      <div className="flex flex-col items-center justify-center gap-3">
        <Button
          onClick={handleClick}
          className="bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white mb-5"
          disabled={loading}
        >
          {loading ? (
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
          ) : (
            <MdAddBox className="text-white mr-2 w-5 h-5" />
          )}
          <p className="text-xl">Create Trigger</p>
        </Button>
        {links.map((item, key) => (
          <Button
            key={key}
            className={`bg-white hover:bg-zinc-100 text-black w-full justify-start rounded-md`}
            asChild
          >
            <Link href={item.href}>
              {item.icon}
              <p className="text-xl">{item.name}</p>
            </Link>
          </Button>
        ))}
      </div>
      <div className="mt-auto">
        <Button
          variant="ghost"
          className="w-full rounded-md justify-start text-red-600 hover:text-red-700 hover:bg-red-100 dark:hover:bg-red-900"
        >
          <MdLogout className="mr-2 w-5 h-5" />
          <p className="text-xl">Logout</p>
        </Button>
      </div>
    </div>
  );
};
