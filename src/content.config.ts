import { defineCollection } from 'astro:content';
import { z } from 'astro/zod';
import { docsLoader } from '@astrojs/starlight/loaders';
import { docsSchema } from '@astrojs/starlight/schema';

export const collections = {
	docs: defineCollection({
		loader: docsLoader(),
		schema: docsSchema({
			extend: z.object({
				/** Last time I re-read / verified this page */
				lastReviewed: z.coerce.date().optional(),
				/** draft = stub not written yet, done = summarised */
				status: z.enum(['draft', 'done']).default('draft'),
			}),
		}),
	}),
};
