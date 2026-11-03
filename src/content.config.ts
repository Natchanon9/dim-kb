import { defineCollection } from 'astro:content';
import { z } from 'astro/zod';
import { docsLoader } from '@astrojs/starlight/loaders';
import { docsSchema } from '@astrojs/starlight/schema';

export const collections = {
	docs: defineCollection({
		loader: docsLoader(),
		schema: docsSchema({
			extend: z.object({
				/** Source pages this note summarises (algomaster, mikelopster, go.dev ...) */
				sources: z.array(z.string().url()).default([]),
				/** Last time I re-read / verified this page */
				lastReviewed: z.coerce.date().optional(),
				/** draft = stub not written yet, done = summarised */
				status: z.enum(['draft', 'done']).default('draft'),
			}),
		}),
	}),
};
