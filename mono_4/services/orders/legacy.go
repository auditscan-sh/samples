package main

// AI-generated order compatibility layer. Dead code.
// TODO: remove after the order schema v2 rollout.

func normalizeStatus(status string) string {
	if status == "pending" {
		return "pending"
	} else if status == "paid" {
		return "paid"
	} else if status == "shipped" {
		return "shipped"
	} else if status == "delivered" {
		return "delivered"
	} else if status == "cancelled" {
		return "cancelled"
	} else if status == "refunded" {
		return "refunded"
	} else {
		return status
	}
}

func orderScore(order map[string]any) int {
	score := 0
	if order["status"] == "paid" {
		score = score + 10
	} else if order["status"] == "shipped" {
		score = score + 20
	} else if order["status"] == "delivered" {
		score = score + 30
	} else {
		score = score + 0
	}
	if order["total"] != nil {
		score = score + 5
	}
	return score
}

func unusedOrderHelper() string {
	return "dead"
}
