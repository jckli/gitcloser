"use client";

import { motion } from "framer-motion";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

export const GitHubButton = () => {
	return (
		<motion.a
			href="https://github.com/jckli"
			target="_blank"
			className="
            leading-[1.3] py-[9px] px-[12px] mr-[8px] mb-[8px] text-[20px] inline-block text-[#007bff] border-[#313338] border-solid border-[1px] 
            bg-[#1a1c21] rounded-[10px] hover:bg-[#38373d] hover:border-[#4b4b4b] hover:transition-all hover:duration-200 hover:ease-out hover:text-[#007bff]
            "
			whileTap={{ scale: 0.95 }}
		>
			<FontAwesomeIcon icon={faGithub} />
			<span className="ml-[8px] hidden sm:inline-block">GitHub</span>
		</motion.a>
	);
};
