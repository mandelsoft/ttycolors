package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings", func() {
	It("handles formatting as feature", func() {
		s := NewSettings(FeatureUnderlineColor)
		Expect(s.IsEnabled()).To(BeFalse())
		s.Enable()
		Expect(s.IsEnabled()).To(BeTrue())
		Expect(s.Features()).To(Equal(FeatureFormatting | FeatureUnderlineColor))
		s.Enable(false)
		Expect(s.IsEnabled()).To(BeFalse())
		Expect(s.Features()).To(Equal(FeatureUnderlineColor))
	})

	It("enables and disables features", func() {
		s := NewSettings(NoFeatures)
		s.EnableFeatures(FeatureFormatting | FeatureUnderlineStyles)
		Expect(s.Features()).To(Equal(FeatureFormatting | FeatureUnderlineStyles))
		s.DisableFeatures(FeatureUnderlineStyles | FeatureUnderlineColor)
		Expect(s.Features()).To(Equal(FeatureFormatting))
		s.SetFeatures(AllFeatures)
		Expect(s.Features()).To(Equal(AllFeatures))
	})

	It("is copied by value", func() {
		s := NewSettings(FeatureFormatting)
		c := s
		c.DisableFeatures(FeatureFormatting)
		Expect(s.IsEnabled()).To(BeTrue())
	})
})
