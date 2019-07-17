package identity

// BaseUserPost defines the request for CreateInactiveUser
type BaseUserPost struct {
	/* public string StudentEmail { get; set; }
	public string FirstName { get; set; }
	public string LastName { get; set; }
	public string OldUserId { get; set; }

	public string Nationality { get; set; }
	public string Avatar { get; set; }
	public int? PrimarySchool { get; set; }
	public int? SchoolRegisteredAt { get; set; } */
}

// BaseVerifyPost defines the request for VerifySchool
type BaseVerifyPost struct {
	/* [Required]
	public Guid UserId { get; set; }

	[Required]
	public int SchoolId { get; set; }

	[Required]
	public string StudentEmail { get; set; } */
}

// ActiveUserPost defines the request for ChangeAuthenticationMethod
type ActiveUserPost struct {
	/* [Required]
	public string Token { get; set; }

	// Activate user with external token providers
	public string Provider { get; set; }
	public string AuthCode { get; set; }
	public string RedirectUrl { get; set; }

	// Activate user with email or password
	public string Email { get; set; }
	public string Password { get; set; }
	public string EmailReturnUrl { get; set; } */
}
