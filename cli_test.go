package headscale

import (
	"time"

	"gopkg.in/check.v1"
)

func (s *Suite) TestRegisterMachine(c *check.C) {
	n, err := h.CreateNamespace("test")
	c.Assert(err, check.IsNil)

	now := time.Now().UTC()

	m := Machine{
		ID:              0,
		MachineKey:      "8ce002a935f8c394e55e78fbbb410576575ff8ec5cfa2e627e4b807f1be15b0e",
		NodeKey:         "bar",
		DiscoKey:        "faa",
		Name:            "testmachine",
		NamespaceID:     n.ID,
		IPAddress:       "10.0.0.1",
		Expiry:          &now,
		RequestedExpiry: &now,
	}
	h.db.Save(&m)

	_, err = h.GetMachine("test", "testmachine")
	c.Assert(err, check.IsNil)

	m2, err := h.RegisterMachine("8ce002a935f8c394e55e78fbbb410576575ff8ec5cfa2e627e4b807f1be15b0e", n.Name)
	c.Assert(err, check.IsNil)
	c.Assert(m2.Registered, check.Equals, true)

	_, err = m2.GetHostInfo()
	c.Assert(err, check.IsNil)
}
