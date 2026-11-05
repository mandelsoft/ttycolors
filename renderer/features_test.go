package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Features", func() {
	It("checks features", func() {
		f := FeatureUnderlineStyles
		Expect(f.Has(FeatureUnderlineStyles)).To(BeTrue())
		Expect(f.Has(FeatureUnderlineColor)).To(BeFalse())
		Expect(f.Has(NoFeatures)).To(BeTrue())
		Expect((f | FeatureUnderlineColor).Has(FeatureUnderlineStyles | FeatureUnderlineColor)).To(BeTrue())
	})

	It("enables and disables features", func() {
		Expect(FeatureFormatting.With(FeatureUnderlineColor, true)).To(Equal(FeatureFormatting | FeatureUnderlineColor))
		Expect(AllFeatures.With(FeatureFormatting, false)).To(Equal(ExtendedFeatures))
	})

	It("provides feature sets", func() {
		Expect(ExtendedFeatures).To(Equal(FeatureUnderlineStyles | FeatureUnderlineColor))
		Expect(AllFeatures.Has(FeatureFormatting | ExtendedFeatures)).To(BeTrue())
		Expect(ExtendedFeatures.Has(FeatureFormatting)).To(BeFalse())
	})
})
